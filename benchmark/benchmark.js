const autocannon = require("autocannon");

const instance = autocannon(
  {
    url: "http://localhost/api/resources/1/book",
    connections: 500,
    amount: 500,
    method: "POST"
  },
  function(err, results) {
    const start = new Date(results.start).getTime();
    const finish = new Date(results.finish).getTime();
    const ts = Math.round(finish - start);
    console.log(`Finished in ${ts}ms`);
    console.log(results["1xx"]);
    console.log(results["2xx"]);
    console.log(results["3xx"]);
    console.log(results["4xx"]);
    console.log(results["5xx"]);
  }
);
