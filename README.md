# eventstore

Simple Knative service persisting Cloud Events to Cloud Firestore collection. Useful in Knative Events demos

## Prerequisites

 * [Knative](https://github.com/knative/docs/blob/master/install) installed
    * Configured [outbound network access (https://github.com/knative/docs/blob/master/serving/outbound-network-access.md)
    * Installed [Knative Eventing](https://github.com/knative/docs/tree/master/eventing) using the `release.yaml` file


## Deployment

Firestore client still requires GCP Project ID to create a client. So, before we can deploy this service to Knative, you will need to update the `GCP_PROJECT_ID` in Now in the `service.yaml` file.

```yaml
    - name: GCP_PROJECT_ID
      value: "enter your project ID here"
```

Once done updating our service manifest (`service.yaml`) you are ready to deploy it.

```shell
kubectl apply -f deployments/service.yaml -n demo
```

The response should be

```shell
service.serving.knative.dev "eventstore" configured
```

To check if the service was deployed successfully you can check the status using `kubectl get pods -n demo` command. The response should look something like this (e.g. Ready `3/3` and Status `Running`).

```shell
NAME                                          READY     STATUS    RESTARTS   AGE
eventstore-0000n-deployment-5645f48b4d-mb24j  3/3       Running   0          10s
```

## Configuration

To make `eventstore` service `cluster.local` so it does not expose externally accessible endpoint but still enable other services to discover it using `eventstore` simply add `serving.knative.dev/visibility: cluster-local` label to `deployments/service.yaml` service manifest

## Disclaimer

This is my personal project and it does not represent my employer. I take no responsibility for issues caused by this code. I do my best to ensure that everything works, but if something goes wrong, my apologies is all you will get.

2019/04/23 13:12:07 Method, POST
2019/04/23 13:12:07 Raw Event: Validation: valid
Context Attributes,
  specversion: 0.2
  type: com.twitter
  source: https://twitter.com/api/1
  id: 1111-2222-3333-4444-5555-6666
  time: 2018-04-05T03:56:24Z
  contenttype: application/json
Data,
  {
    "text": "RT @PostGradProblem: In preparation for the NFL lockout, I will be spending twice as much time analyzing my fantasy baseball team during ...",
    "favorited": false,
    "source": "<a href=\"http://twitter.com/\" rel=\"nofollow\">Twitter for iPhone</a>",
    "id_str": "54691802283900928",
    "entities": {
      "user_mentions": [
        {
          "indices": [
            3,
            19
          ],
          "screen_name": "PostGradProblem",
          "id_str": "271572434",
          "name": "PostGradProblems",
          "id": 271572434
        }
      ],
      "urls": [],
      "hashtags": []
    },
    "retweeted": false,
    "retweet_count": 4,
    "created_at": "Sun Apr 03 23:48:36 +0000 2011",
    "retweeted_status": {
      "text": "In preparation for the NFL lockout, I will be spending twice as much time analyzing my fantasy baseball team during company time. #PGP",
      "truncated": false,
      "favorited": false,
      "source": "<a href=\"http://www.hootsuite.com\" rel=\"nofollow\">HootSuite</a>",
      "id_str": "54640519019642881",
      "entities": {
        "user_mentions": [],
        "urls": [],
        "hashtags": [
          {
            "text": "PGP",
            "indices": [
              130,
              134
            ]
          }
        ]
      },
      "retweeted": false,
      "place": null,
      "retweet_count": 4,
      "created_at": "Sun Apr 03 20:24:49 +0000 2011",
      "user": {
        "notifications": null,
        "profile_use_background_image": true,
        "statuses_count": 31,
        "profile_background_color": "C0DEED",
        "followers_count": 3066,
        "profile_image_url": "http://a2.twimg.com/profile_images/1285770264/PGP_normal.jpg",
        "listed_count": 6,
        "profile_background_image_url": "http://a3.twimg.com/a/1301071706/images/themes/theme1/bg.png",
        "description": "",
        "screen_name": "PostGradProblem",
        "default_profile": true,
        "verified": false,
        "time_zone": null,
        "profile_text_color": "333333",
        "is_translator": false,
        "profile_sidebar_fill_color": "DDEEF6",
        "location": "",
        "id_str": "271572434",
        "default_profile_image": false,
        "profile_background_tile": false,
        "lang": "en",
        "friends_count": 21,
        "protected": false,
        "favourites_count": 0,
        "created_at": "Thu Mar 24 19:45:44 +0000 2011",
        "profile_link_color": "0084B4",
        "name": "PostGradProblems",
        "show_all_inline_media": false,
        "follow_request_sent": null,
        "geo_enabled": false,
        "profile_sidebar_border_color": "C0DEED",
        "url": null,
        "id": 271572434,
        "contributors_enabled": false,
        "following": null,
        "utc_offset": null
      },
      "id": 54640519019642880,
      "coordinates": null,
      "geo": null
    },
    "user": {
      "notifications": null,
      "profile_use_background_image": true,
      "statuses_count": 351,
      "profile_background_color": "C0DEED",
      "followers_count": 48,
      "profile_image_url": "http://a1.twimg.com/profile_images/455128973/gCsVUnofNqqyd6tdOGevROvko1_500_normal.jpg",
      "listed_count": 0,
      "profile_background_image_url": "http://a3.twimg.com/a/1300479984/images/themes/theme1/bg.png",
      "description": "watcha doin in my waters?",
      "screen_name": "OldGREG85",
      "default_profile": true,
      "verified": false,
      "time_zone": "Hawaii",
      "profile_text_color": "333333",
      "is_translator": false,
      "profile_sidebar_fill_color": "DDEEF6",
      "location": "Texas",
      "id_str": "80177619",
      "default_profile_image": false,
      "profile_background_tile": false,
      "lang": "en",
      "friends_count": 81,
      "protected": false,
      "favourites_count": 0,
      "created_at": "Tue Oct 06 01:13:17 +0000 2009",
      "profile_link_color": "0084B4",
      "name": "GG",
      "show_all_inline_media": false,
      "follow_request_sent": null,
      "geo_enabled": false,
      "profile_sidebar_border_color": "C0DEED",
      "url": null,
      "id": 80177619,
      "contributors_enabled": false,
      "following": null,
      "utc_offset": -36000
    },
    "id": 54691802283900930,
    "coordinates": null,
    "geo": null
  }
2019/04/23 13:12:07 Text to score: "RT @PostGradProblem: In preparation for the NFL lockout, I will be spending twice as much time analyzing my fantasy baseball team during ..."
2019/04/23 13:12:07 Score: map[magnitude:0.2 score:0.2]
2019/04/23 13:12:07 Processed Event: Validation: valid
Context Attributes,
  specversion: 0.2
  type: com.twitter.scored
  source: https://twitter.com/api/1
  id: 1111-2222-3333-4444-5555-6666
  time: 2018-04-05T03:56:24Z
  contenttype: application/json
Extensions,
  sentiment: map[magnitude:0.2 score:0.2]
Data,
  {
    "text": "RT @PostGradProblem: In preparation for the NFL lockout, I will be spending twice as much time analyzing my fantasy baseball team during ...",
    "favorited": false,
    "source": "<a href=\"http://twitter.com/\" rel=\"nofollow\">Twitter for iPhone</a>",
    "id_str": "54691802283900928",
    "entities": {
      "user_mentions": [
        {
          "indices": [
            3,
            19
          ],
          "screen_name": "PostGradProblem",
          "id_str": "271572434",
          "name": "PostGradProblems",
          "id": 271572434
        }
      ],
      "urls": [],
      "hashtags": []
    },
    "retweeted": false,
    "retweet_count": 4,
    "created_at": "Sun Apr 03 23:48:36 +0000 2011",
    "retweeted_status": {
      "text": "In preparation for the NFL lockout, I will be spending twice as much time analyzing my fantasy baseball team during company time. #PGP",
      "truncated": false,
      "favorited": false,
      "source": "<a href=\"http://www.hootsuite.com\" rel=\"nofollow\">HootSuite</a>",
      "id_str": "54640519019642881",
      "entities": {
        "user_mentions": [],
        "urls": [],
        "hashtags": [
          {
            "text": "PGP",
            "indices": [
              130,
              134
            ]
          }
        ]
      },
      "retweeted": false,
      "place": null,
      "retweet_count": 4,
      "created_at": "Sun Apr 03 20:24:49 +0000 2011",
      "user": {
        "notifications": null,
        "profile_use_background_image": true,
        "statuses_count": 31,
        "profile_background_color": "C0DEED",
        "followers_count": 3066,
        "profile_image_url": "http://a2.twimg.com/profile_images/1285770264/PGP_normal.jpg",
        "listed_count": 6,
        "profile_background_image_url": "http://a3.twimg.com/a/1301071706/images/themes/theme1/bg.png",
        "description": "",
        "screen_name": "PostGradProblem",
        "default_profile": true,
        "verified": false,
        "time_zone": null,
        "profile_text_color": "333333",
        "is_translator": false,
        "profile_sidebar_fill_color": "DDEEF6",
        "location": "",
        "id_str": "271572434",
        "default_profile_image": false,
        "profile_background_tile": false,
        "lang": "en",
        "friends_count": 21,
        "protected": false,
        "favourites_count": 0,
        "created_at": "Thu Mar 24 19:45:44 +0000 2011",
        "profile_link_color": "0084B4",
        "name": "PostGradProblems",
        "show_all_inline_media": false,
        "follow_request_sent": null,
        "geo_enabled": false,
        "profile_sidebar_border_color": "C0DEED",
        "url": null,
        "id": 271572434,
        "contributors_enabled": false,
        "following": null,
        "utc_offset": null
      },
      "id": 54640519019642880,
      "coordinates": null,
      "geo": null
    },
    "user": {
      "notifications": null,
      "profile_use_background_image": true,
      "statuses_count": 351,
      "profile_background_color": "C0DEED",
      "followers_count": 48,
      "profile_image_url": "http://a1.twimg.com/profile_images/455128973/gCsVUnofNqqyd6tdOGevROvko1_500_normal.jpg",
      "listed_count": 0,
      "profile_background_image_url": "http://a3.twimg.com/a/1300479984/images/themes/theme1/bg.png",
      "description": "watcha doin in my waters?",
      "screen_name": "OldGREG85",
      "default_profile": true,
      "verified": false,
      "time_zone": "Hawaii",
      "profile_text_color": "333333",
      "is_translator": false,
      "profile_sidebar_fill_color": "DDEEF6",
      "location": "Texas",
      "id_str": "80177619",
      "default_profile_image": false,
      "profile_background_tile": false,
      "lang": "en",
      "friends_count": 81,
      "protected": false,
      "favourites_count": 0,
      "created_at": "Tue Oct 06 01:13:17 +0000 2009",
      "profile_link_color": "0084B4",
      "name": "GG",
      "show_all_inline_media": false,
      "follow_request_sent": null,
      "geo_enabled": false,
      "profile_sidebar_border_color": "C0DEED",
      "url": null,
      "id": 80177619,
      "contributors_enabled": false,
      "following": null,
      "utc_offset": -36000
    },
    "id": 54691802283900930,
    "coordinates": null,
    "geo": null
  }
^Csignal: interrupt