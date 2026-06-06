angular.module('opendocking.docker').controller('NodeAvailabilitySelectController', [
  function NodeAvailabilitySelectController() {
    this.onChange = onChange;

    function onChange() {
      this.onSave({ availability: this.availability });
    }
  },
]);
