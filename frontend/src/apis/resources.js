import axios from "axios";

function getFullUrl(url) {
  return `${process.env.REACT_APP_API_URL}${url}`;
}

export default {
  getAll: async function(params) {
    const url = getFullUrl("/resources");
    const result = await axios.get(url);
    return result.data;
  },
  book: async function(resourceId) {
    const url = getFullUrl(`/resources/${resourceId}/book`);
    const result = await axios.post(url);
    return result.data;
  },
  resetAll: async function() {
    const url = getFullUrl("/resources/restore");
    const result = await axios.post(url);
    return result.data;
  }
};
