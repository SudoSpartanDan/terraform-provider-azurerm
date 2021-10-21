package features

func Default() UserFeatures {
	return UserFeatures{
		// NOTE: ensure all nested objects are fully populated
		ApiManagement: ApiManagementFeatures{
			PurgeSoftDeleteOnDestroy: false,
		},
		CognitiveAccount: CognitiveAccountFeatures{
			PurgeSoftDeleteOnDestroy: true,
		},
		KeyVault: KeyVaultFeatures{
			PurgeSoftDeleteOnDestroy:         true,
			RecoverSoftDeletedKeyVaults:      true,
			RecoverSoftDeletedKeys:           true,
			PurgeSoftDeletedKeysOnDestroy:    true,
			RecoverSoftDeletedCerts:          true,
			PurgeSoftDeletedCertsOnDestroy:   true,
			RecoverSoftDeletedSecrets:        true,
			PurgeSoftDeletedSecretsOnDestroy: true,
		},
		LogAnalyticsWorkspace: LogAnalyticsWorkspaceFeatures{
			PermanentlyDeleteOnDestroy: false,
		},
		Network: NetworkFeatures{
			RelaxedLocking: false,
		},
		ResourceGroup: ResourceGroupFeatures{
			PreventDeletionIfContainsResources: false,
		},
		TemplateDeployment: TemplateDeploymentFeatures{
			DeleteNestedItemsDuringDeletion: true,
		},
		VirtualMachine: VirtualMachineFeatures{
			DeleteOSDiskOnDeletion:     true,
			GracefulShutdown:           false,
			SkipShutdownAndForceDelete: false,
		},
		VirtualMachineScaleSet: VirtualMachineScaleSetFeatures{
			ForceDelete:               false,
			RollInstancesWhenRequired: true,
		},
	}
}
