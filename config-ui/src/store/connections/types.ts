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

export enum ConnectionStatusEnum {
  ONLINE = 'online',
  OFFLINE = 'offline',
  TESTING = 'testing',
  NULL = 'null',
}

export type ConnectionItemType = {
  unique: string;
  status: ConnectionStatusEnum;
  plugin: string;
  id: ID;
  name: string;
  icon: string;
  entities: string[];
  transformationType: 'none' | 'for-connection' | 'for-scope';
  endpoint: string;
  proxy: string;
  token?: string;
  username?: string;
  password?: string;
  spreadsheetID?: string;
  firstValue?: string;
  lastValue?: string;
  authMethod?: string;
};
