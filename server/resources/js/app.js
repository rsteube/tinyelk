requirejs.config({
  baseUrl: 'js/lib',
  paths: {
    app: '../app',
  }
});

require(['app/view/Main'], function(Main) {
  m.mount(document.body, Main)
});
