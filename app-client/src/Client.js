/* eslint-disable no-undef */
function translate(value) {
  return fetch(`tags`, {
    accept: "application/json"
  })
    .then(checkStatus)
    .then(parseJSON)
}

function search(query) {
  return fetch(`filteredtags/?key=${query}`, {
    accept: "application/json"
  }) 
    .then(checkStatus)
    .then(parseJSON)
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

const Client = { translate, search };
export default Client;
