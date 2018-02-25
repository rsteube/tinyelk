define(function() {
  var Log = {
    list: [],
    loadList: function() {
      return m.request({
          method: 'GET',
          url: '//localhost:7318/jq?q=.',
        })
        .then(function(result) {
          Log.list = result
        })
    }
  }
  return Log
});
