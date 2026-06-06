angular.module('opendocking.kubernetes').component('kubernetesConfigurationData', {
  templateUrl: './kubernetesConfigurationData.html',
  controller: 'KubernetesConfigurationDataController',
  bindings: {
    formValues: '=',
    isDockerConfig: '=',
    onChangeValidation: '&',
    isValid: '=',
    isCreation: '=',
    isEditorDirty: '=',
    type: '<',
  },
});
