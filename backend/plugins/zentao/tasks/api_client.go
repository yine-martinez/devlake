/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tasks

import (
	"net/http"
	"strconv"
	"time"

	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/zentao/models"
)

func NewZentaoApiClient(taskCtx plugin.TaskContext, connection *models.ZentaoConnection) (*api.ApiAsyncClient, error) {

	authApiClient, err := api.NewApiClient(taskCtx.GetContext(), connection.Endpoint, nil, 0, connection.Proxy, taskCtx)
	if err != nil {
		return nil, err
	}
	// request for access token
	tokenReqBody := &models.ApiAccessTokenRequest{
		Account:  connection.Username,
		Password: connection.Password,
	}
	tokenRes, err := authApiClient.Post("/tokens", nil, tokenReqBody, nil)
	if err != nil {
		return nil, err
	}
	tokenResBody := &models.ApiAccessTokenResponse{}
	err = api.UnmarshalResponse(tokenRes, tokenResBody)
	apiClient, err := api.NewApiClientFromConnection(taskCtx.GetContext(), taskCtx, connection)
	if err != nil {
		return nil, err
	}
	// create rate limit calculator
	rateLimiter := &api.ApiRateLimitCalculator{
		UserRateLimitPerHour: connection.RateLimitPerHour,
		DynamicRateLimit: func(res *http.Response) (int, time.Duration, errors.Error) {
			rateLimitHeader := res.Header.Get("RateLimit-Limit")
			if rateLimitHeader == "" {
				// use default
				return 0, 0, nil
			}
			rateLimit, err := strconv.Atoi(rateLimitHeader)
			if err != nil {
				return 0, 0, errors.Default.Wrap(err, "failed to parse RateLimit-Limit header: %w")
			}
			// seems like {{ .plugin-ame }} rate limit is on minute basis
			return rateLimit, 1 * time.Minute, nil
		},
	}
	asyncApiClient, err := api.CreateAsyncApiClient(
		taskCtx,
		apiClient,
		rateLimiter,
	)
	if err != nil {
		return nil, err
	}
	return asyncApiClient, nil
}

type ZentaoPagination struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
}
