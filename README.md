# pubsub-twitter-event-viewer [![Build Status](https://travis-ci.org/mchmarny/pubsub-twitter-event-viewer.svg?branch=master)](https://travis-ci.org/mchmarny/pubsub-twitter-event-viewer)

Web app to view tweets published to Google PubSub by [twitter-to-pubsub-event-pump](https://github.com/mchmarny/twitter-to-pubsub-event-pump)



## Setup

### GCP

If you don't already have GCP account, you can run this entire app using the Google Cloud Platform (GCP) [free tier](https://cloud.google.com/free/). Once you create project, you will need to pass the `GCLOUD_PROJECT` argument on each execution or you can define it as environment variables like this:

```shell 
export GCLOUD_PROJECT="YOUR_PROJECT_NAME"
```

### GCP CLI

If you don't already have `gcloud`, you can find instructions on how to download and install the GCP SDK [here](https://cloud.google.com/sdk/)


#### Service Account 

You will need to set up GCP authentication using service account. You can find instructions how to do this [here](https://cloud.google.com/video-intelligence/docs/common/auth#set_up_a_service_account). After you download your service account file you will need to define

```shell 
export GOOGLE_APPLICATION_CREDENTIALS=<path_to_service_account_file>
```

#### Create PubSub subscription

You can create GCP PubSub topic `gcloud` by executing the following command:

```shell 
gcloud pubsub subscriptions create tweets-sub --topic tweets
```

## Build

To first build the app you can execute first `make dep` which will assure your environment has the necessary dependencies. Alternatively you can restore the app dependencies 

```shell
go get github.com/tools/godep
godep restore
```

and then `make build` to build the app or alternatively you can run the build command directly

```shell
go build -v -o bin/tview
```

## Run 

You can run the viewer app (assuming it's already built) using the following command: `make run`. Alternatively you can run it by executing the `tview` binary built in previous step. 

```shell
bin/tview
```

Once the server is started you can navigate to [http://127.0.0.1:8080]() to view the published tweets. 
