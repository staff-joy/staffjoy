import { MILISECONDS_TO_SECONDS } from './constants/constants';
import {
  ENV_NAME_DEVELOPMENT,
  ENV_NAME_STAGING,
  ENV_NAME_PRODUCTION,
  DEVELOPMENT_APEX,
  STAGING_APEX,
  PRODUCTION_APEX,
  HTTP_PREFIX,
  HTTPS_PREFIX,
  getRefetchInterval,
} from './constants/config';

export function emptyPromise(val = null) {
  /* creates an empty promise for cases when data doesn't need to be fetched */
  return new Promise((resolve) => { resolve(val); });
}

export function timestampExpired(timestamp, endpoint = 'DEFAULT') {
  /*
    input: timestamp and a (str) endpoint name
    output: true if the timestamp has elapsed longer than the endpoint allows
  */
  const timeDiff = (Date.now() - timestamp) / MILISECONDS_TO_SECONDS;
  return timeDiff > getRefetchInterval(endpoint);
}

export function detectEnvironment() {
  let env = ENV_NAME_DEVELOPMENT;
  const url = window.location.href.toLowerCase();
  const domain = url.split('/')[2];

  if (domain.endsWith(PRODUCTION_APEX)) {
    env = ENV_NAME_PRODUCTION;
  } else if (domain.endsWith(STAGING_APEX)) {
    env = ENV_NAME_STAGING;
  }

  return env;
}

export function routeToMicroservice(service, path = '') {
  const devRoute = `${HTTP_PREFIX}${service}${DEVELOPMENT_APEX}${path}`;

  switch (detectEnvironment()) {
    case ENV_NAME_DEVELOPMENT:
      return devRoute;

    case ENV_NAME_STAGING:
      return `${HTTPS_PREFIX}${service}${STAGING_APEX}${path}`;

    case ENV_NAME_PRODUCTION:
      return `${HTTPS_PREFIX}${service}${PRODUCTION_APEX}${path}`;

    default:
      return devRoute;
  }
}

export function checkStatus(response) {
  if (response.status >= 200 && response.status < 300) {
    return response;
  }

  const error = new Error(response.statusText);
  error.response = response;
  throw error;
}

export function parseJSON(response) {
  return response.json();
}
