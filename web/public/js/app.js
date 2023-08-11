'use strict';

(function () {
  function init() {
    var router = new Router([
      new Route('list-vehicles', 'vehicle/list.html', true),
    ]);
  }
  init();
})();
