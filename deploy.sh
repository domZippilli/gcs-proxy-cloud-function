#!/usr/bin/env bash
# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

function usage(){
    echo >&2
    echo "Usage: $0 BUCKET_NAME" >&2
    echo "Deploys a Cloud Functions proxy of a GCS bucket." >&2
    echo >&2
}
BUCKET_NAME=${1?$(usage)}

gcloud functions deploy GCS-"${BUCKET_NAME}" \
--entry-point=ProxyGCS \
--runtime go113 \
--memory=128MB \
--set-env-vars BUCKET_NAME="${BUCKET_NAME}" \
--trigger-http \
--allow-unauthenticated \
--security-level=secure-always