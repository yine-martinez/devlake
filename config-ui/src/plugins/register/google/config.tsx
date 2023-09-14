/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

import React from 'react';

import type { PluginConfigType } from '@/plugins';
import { PluginType } from '@/plugins';

import Icon from './assets/icon.svg';

import { SpreadsheetID, FirstValue, LastValue } from './connection-fields';

export const GoogleConfig: PluginConfigType = {
  type: PluginType.Connection,
  plugin: 'google',
  name: 'google',
  icon: Icon,
  sort: 101,
  connection: {
    docLink: 'https://devlake.apache.org/docs/Configuration/GitHub',
    initialValues: {
      endpoint: 'https://sheets.googleapis.com/v4',
    },
    fields: [
      'name',
      'token',
      'endpoint',
      ({ initialValues, values, errors, setValues, setErrors }: any) => (
        <SpreadsheetID
          name="spreadsheetID"
          key="spreadsheetID"
          initialValue={initialValues.spreadsheetID ?? ''}
          value={values.spreadsheetID ?? ''}
          error={errors.spreadsheetID ?? ''}
          setValue={(value) => setValues({ spreadsheetID: value })}
          setError={(value) => setErrors({ spreadsheetID: value })}
        />
      ),
      ({ initialValues, values, setValues, setErrors }: any) => (
        <FirstValue
          initialValue={initialValues.firstValue ?? ''}
          value={values.firstValue ?? ''}
          setValue={(value) => setValues({ firstValue: value })}
          setError={(value) => setErrors({ firstValue: value })}
        />
      ),
      ({ initialValues, values, errors, setValues, setErrors }: any) => (
        <LastValue
          key="LastValue"
          initialValue={initialValues.lastValue ?? ''}
          value={values.lastValue ?? ''}
          error={errors.lastValue ?? ''}
          setValue={(value) => setValues({ lastValue: value })}
          setError={(value) => setErrors({ lastValue: value })}
        />
      ),
    ],
  },
  entities: [],
  transformation: {},
};
