# mailing-list

[![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](LICENSE)
[![CircleCI](https://circleci.com/gh/tiny-pilot/mailing-list.svg?style=svg&circle-token=d4e7bed824bf9cfd1baff5c84bb9eee541fcfe23)](https://circleci.com/gh/tiny-pilot/mailing-list)

Signup handler for TinyPilot Mailing List.

## Details

This is a wrapper for the [EmailOctopus API](https://emailoctopus.com/api-documentation) that exposes a user-facing REST interface for signing up for a mailing list. It runs as a [Google Cloud Function](https://cloud.google.com/functions/docs/concepts/exec).

## Who can use it

You can use this project if you want to implement your own mailing list signup frontend for EmailOctopus rather than use one of their pre-generated forms.

## Deployment

This project is configured to automatically deploy from CircleCI. If you use a difference CI environment or want to deploy manually, you can copy the contents of [`.circleci/config.yml`](./circleci/config.yml). The relevant environment variables are:

* `CLIENT_SECRET`: Your Google Cloud Project client secret JSON file, encoded as a base64 string. The user associated with the client secret must have the following GCP roles: `appengine.appAdmin`, `cloudbuild.builds.builder`, `cloudbuild.builds.editor`, `cloudfunctions.admin`, `iam.serviceAccountUser`, `run.invoker`, `storage.admin`, `storage.objectAdmin`
* `EMAIL_OCTOPUS_API_KEY`: Your API key from EmailOctopus.
* `EMAIL_OCTOPUS_LIST_ID`: The ID of the list
* `GCLOUD_PROJECT`: The name of your Google Cloud Platform project
* `GO_RUNTIME`: Which Google Cloud Function to use (go113 recommended)
* `FUNCTION_NAME`: Name to give your deployed Google Cloud Function

## Integration

The following JavaScript snippet demonstrates a sample

```javascript
/* Add a subscriber to the EmailOctopus mailing list

@param {string} email: The email address of the user to add.
@param {string} ninja: A honeypot field. If it's anything but an empty
  string, the server will report success but discard the signup. It's
  meant to prevent bot signups. It should represent the value in an input
  field that's hidden from real users but that automated bots would fill
  with junk.
*/
function subscribe(email, ninja) {
  // Replace this with the URL where you host your Google Cloud Function.
  const functionUrl = 'https://cloud-fns.tinypilotkvm.com/email-signup';

  return fetch(functionUrl, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({email, ninja}),
  }).then((response) => {
    if (!response.ok) {
      return response.text().then((error) => {
        return Promise.reject(new Error(error));
      });
    }
    return Promise.resolve();
  });
```
