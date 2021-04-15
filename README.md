# gcs-proxy-cloud-function

A Cloud Function to proxy a GCS bucket with an HTTP Cloud Function. Useful for conditional serving logic, transcoding, certain security features, etc.

Many use cases can be satisfied by using [built-in static website hosting for GCS with a Cloud Load Balancer](https://cloud.google.com/storage/docs/hosting-static-website). If there are limitations of that feature that are blocking you, this proxy approach might be for you.

However, consider using [gcs-proxy-cloud-run](https://github.com/domZippilli/gcs-proxy-cloud-run) instead. This allows the proxy to stream
responses to the client, whereas this proxy will have to "store and forward" since Cloud Function responses can only be sent once the function
closes.

## Deployment

As a prerequisite, [enable the Cloud Build API](https://console.cloud.google.com/apis/library/cloudbuild.googleapis.com) for your project.

Also, if you haven't done so, ensure `gcloud` is using the correct credentials. Usually, a combination of `gcloud auth login`, `gcloud config set project`, and optionally `gcloud auth revoke` when you are finished will do the job.

Then, simply run `deploy.sh` with the bucket you want to proxy as the first argument:

```shell
./deploy.sh mystaticcontent
```

## Copyright

```text
Copyright 2021 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
