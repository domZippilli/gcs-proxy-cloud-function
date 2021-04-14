# gcs-proxy-cloud-function

A Cloud Function to proxy a GCS bucket with an HTTP Cloud Function. Useful for conditional serving logic, transcoding, certain security features, etc.

Many use cases can be satisfied by using [built-in static website hosting for GCS with a Cloud Load Balancer](https://cloud.google.com/storage/docs/hosting-static-website). If there are limitations of that feature that are blocking you, this proxy approach might be for you. 

## Deployment

As a prerequisite, [enable the Cloud Build API](https://console.cloud.google.com/apis/library/cloudbuild.googleapis.com) for your project.

Also, if you haven't done so, ensure `gcloud` is using the correct credentials. Usually, a combination of `gcloud auth login`, `gcloud config set project`, and optionally `gcloud auth revoke` when you are finished will do the job.

Then, simply run `deploy.sh` with the bucket you want to proxy as the first argument:

```shell
./deploy.sh mystaticcontent
```
