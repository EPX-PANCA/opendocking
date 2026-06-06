angular.module('opendocking.app').component('productList', {
  templateUrl: './productList.html',
  bindings: {
    titleText: '@',
    products: '<',
    goTo: '<',
  },
});
