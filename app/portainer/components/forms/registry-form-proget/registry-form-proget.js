class controller {
  $postLink() {
    this.registryFormProGet.registry_name.$validators.used = (modelValue) => !this.nameIsUsed(modelValue);
  }
}

angular.module('opendocking.app').component('registryFormProget', {
  templateUrl: './registry-form-proget.html',
  bindings: {
    model: '=',
    formAction: '<',
    formActionLabel: '@',
    actionInProgress: '<',
    nameIsUsed: '<',
  },
  controller,
});
