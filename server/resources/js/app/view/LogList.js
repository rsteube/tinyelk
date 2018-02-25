define(['app/models/Log'], function(Log) {
  return {
    activeEntry: "",
    oninit: Log.loadList,
    view: function(vnode) {
      var color = {
        'trace': '',
        'debug': '',
        'info': 'has-text-info',
        'warn': 'has-text-warning',
        'error': 'has-text-danger'
      }

      return m('table', {
          class: 'table is-hoverable is-fullwidth is-size-7'
        },
        [
        m('thead',
          m('tr',
          [m('th'), m('th', '_timestamp'), m('th', '_msg')]
          )
        ),
        m('tbody',
          Log.list.map(function(log) {
              return m('tr', {onclick: function(){vnode.state.activeEntry = (vnode.state.activeEntry == log._timestamp ? "" : log._timestamp); m.redraw()}}, [
              m('td', {
                class: 'has-text-centered has-text-weight-bold ' + color[log._level.toLowerCase()]
              }, log._level[0]), m('td', log._timestamp), 
              
              m('td', [log.msg,
              m('table', {id: 'logDetail', class: 'table is-hoverable is-fullwidth ' + (vnode.state.activeEntry == log._timestamp? '' : 'is-hidden')},
                Object.entries(log).map(function([key, value]) {
                  return m('tr', [
                    m('td', key),
                    m('td', {class: ''+ color[value.toLowerCase()]}, value)
                    ])
                })
              )])
            
            
            ])
          })
        )]
      )
    }
  }
});
