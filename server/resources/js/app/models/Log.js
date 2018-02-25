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
    },
    visible: new Set(['trace', 'debug', 'info', 'warn', 'error']),
    toggle: function(level) {
      Log.visible.has(level)? Log.visible.delete(level) : Log.visible.add(level)
    }
  }
  return Log
});
