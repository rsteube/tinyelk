define(['c3', 'd3', 'app/models/Log'], function(c3, d3, Log) {
  return {
    view: function() {
      return m('div', {
        id: 'rangeSelection'
      })
    },
    oncreate: function() {
      var chart = c3.generate({
        bindto: '#rangeSelection',
        color: {
          pattern: [
            '#1abc9c', // trace
            '#343c3d', // debug
            '#209cee', // info
            '#f1b70e', // warn
            '#e74c3c', // error
          ]
        },
        data: {
          columns: [
            ['trace', 30, 200, 100, 400, 150, 250],
            ['debug', 2, 2, 10, 40, 15, 25],
            ['info', 30, 20, 10, 40, 15, 25],
            ['warn', 50, 20, 10, 4, 15, 5],
            ['error', 10, 2, 3, 40, 15, 25]
          ],
        },
        subchart: {
          show: true
        },
        legend: {
          item: {
            onclick: function(id) { chart.toggle(id); Log.toggle(id); m.redraw() }
          }
        }
      })
    }
  }
});
