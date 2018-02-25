define(['app/view/LogList', 'app/view/RangeSelection'], function(LogList, RangeSelection) {
  return {
    view: function() {
      return m('section', {
          class: 'section'
        },[
          m('container', {class: 'container'},
          m(RangeSelection)),
        m('container', {
          class: 'container'
        }, m(LogList))
        ]
      )
    }
  }
});
