#!/usr/bin/env bash
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