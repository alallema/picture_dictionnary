/* eslint-disable no-undef */
function translate(value, cb) {
  return fetch(`tags`, {
    accept: "application/json"
  })
    .then(checkStatus)
    .then(parseJSON)
    .then(function(data) { return getTags(value, data); })
    .then(cb);
}

function search(query, cb) {
  return fetch(`picture/${query}`, {
    accept: "application/json"
  }) 
    .then(checkStatus)
    .then(parseJSON)
    .then(cb);
}

function checkStatus(response) {
  if (response.status >= 200 && response.status < 300) {
    return response;
  }
  const error = new Error(`HTTP Error ${response.statusText}`);
  error.status = response.statusText;
  error.response = response;
  console.log(error); // eslint-disable-line no-console
  throw error;
}

function parseJSON(response) {
  return response.json();
}

function getTags(value, response) {
  var results = [];

  var arr = response.result
  for(var i=0; i<arr.length; i++) {
      if (arr[i].title.toLowerCase().search(value.toLowerCase())!==-1) {
        results.push(arr[i]);
    }
  }
  return results;
}

const Client = { translate, search };
export default Client;
