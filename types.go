package main

type unity_meta struct {
	FileFormatVersion int    `yaml:"fileFormatVersion"`
	GUID              string `yaml:"guid"`
	ModelImporter     struct {
		SerializedVersion     int `yaml:"serializedVersion"`
		InternalIDToNameTable []struct {
			First struct {
				Num74 int64 `yaml:"74"`
			} `yaml:"first"`
			Second string `yaml:"second"`
		} `yaml:"internalIDToNameTable"`
		ExternalObjects struct {
		} `yaml:"externalObjects"`
		Materials struct {
			MaterialImportMode int `yaml:"materialImportMode"`
			MaterialName       int `yaml:"materialName"`
			MaterialSearch     int `yaml:"materialSearch"`
			MaterialLocation   int `yaml:"materialLocation"`
		} `yaml:"materials"`
		Animations struct {
			LegacyGenerateAnimations       int           `yaml:"legacyGenerateAnimations"`
			BakeSimulation                 int           `yaml:"bakeSimulation"`
			ResampleCurves                 int           `yaml:"resampleCurves"`
			OptimizeGameObjects            int           `yaml:"optimizeGameObjects"`
			RemoveConstantScaleCurves      int           `yaml:"removeConstantScaleCurves"`
			MotionNodeName                 interface{}   `yaml:"motionNodeName"`
			RigImportErrors                interface{}   `yaml:"rigImportErrors"`
			RigImportWarnings              interface{}   `yaml:"rigImportWarnings"`
			AnimationImportErrors          interface{}   `yaml:"animationImportErrors"`
			AnimationImportWarnings        interface{}   `yaml:"animationImportWarnings"`
			AnimationRetargetingWarnings   interface{}   `yaml:"animationRetargetingWarnings"`
			AnimationDoRetargetingWarnings int           `yaml:"animationDoRetargetingWarnings"`
			ImportAnimatedCustomProperties int           `yaml:"importAnimatedCustomProperties"`
			ImportConstraints              int           `yaml:"importConstraints"`
			AnimationCompression           int           `yaml:"animationCompression"`
			AnimationRotationError         float64       `yaml:"animationRotationError"`
			AnimationPositionError         float64       `yaml:"animationPositionError"`
			AnimationScaleError            float64       `yaml:"animationScaleError"`
			AnimationWrapMode              int           `yaml:"animationWrapMode"`
			ExtraExposedTransformPaths     []interface{} `yaml:"extraExposedTransformPaths"`
			ExtraUserProperties            []interface{} `yaml:"extraUserProperties"`
			ClipAnimations                 []interface{} `yaml:"clipAnimations"`
			IsReadable                     int           `yaml:"isReadable"`
		} `yaml:"animations"`
		Meshes struct {
			LODScreenPercentages             []interface{} `yaml:"lODScreenPercentages"`
			GlobalScale                      int           `yaml:"globalScale"`
			MeshCompression                  int           `yaml:"meshCompression"`
			AddColliders                     int           `yaml:"addColliders"`
			UseSRGBMaterialColor             int           `yaml:"useSRGBMaterialColor"`
			SortHierarchyByName              int           `yaml:"sortHierarchyByName"`
			ImportVisibility                 int           `yaml:"importVisibility"`
			ImportBlendShapes                int           `yaml:"importBlendShapes"`
			ImportCameras                    int           `yaml:"importCameras"`
			ImportLights                     int           `yaml:"importLights"`
			NodeNameCollisionStrategy        int           `yaml:"nodeNameCollisionStrategy"`
			FileIdsGeneration                int           `yaml:"fileIdsGeneration"`
			SwapUVChannels                   int           `yaml:"swapUVChannels"`
			GenerateSecondaryUV              int           `yaml:"generateSecondaryUV"`
			UseFileUnits                     int           `yaml:"useFileUnits"`
			KeepQuads                        int           `yaml:"keepQuads"`
			WeldVertices                     int           `yaml:"weldVertices"`
			BakeAxisConversion               int           `yaml:"bakeAxisConversion"`
			PreserveHierarchy                int           `yaml:"preserveHierarchy"`
			SkinWeightsMode                  int           `yaml:"skinWeightsMode"`
			MaxBonesPerVertex                int           `yaml:"maxBonesPerVertex"`
			MinBoneWeight                    float64       `yaml:"minBoneWeight"`
			OptimizeBones                    int           `yaml:"optimizeBones"`
			MeshOptimizationFlags            int           `yaml:"meshOptimizationFlags"`
			IndexFormat                      int           `yaml:"indexFormat"`
			SecondaryUVAngleDistortion       int           `yaml:"secondaryUVAngleDistortion"`
			SecondaryUVAreaDistortion        float64       `yaml:"secondaryUVAreaDistortion"`
			SecondaryUVHardAngle             int           `yaml:"secondaryUVHardAngle"`
			SecondaryUVMarginMethod          int           `yaml:"secondaryUVMarginMethod"`
			SecondaryUVMinLightmapResolution int           `yaml:"secondaryUVMinLightmapResolution"`
			SecondaryUVMinObjectScale        int           `yaml:"secondaryUVMinObjectScale"`
			SecondaryUVPackMargin            int           `yaml:"secondaryUVPackMargin"`
			UseFileScale                     int           `yaml:"useFileScale"`
		} `yaml:"meshes"`
		TangentSpace struct {
			NormalSmoothAngle                                                int `yaml:"normalSmoothAngle"`
			NormalImportMode                                                 int `yaml:"normalImportMode"`
			TangentImportMode                                                int `yaml:"tangentImportMode"`
			NormalCalculationMode                                            int `yaml:"normalCalculationMode"`
			LegacyComputeAllNormalsFromSmoothingGroupsWhenMeshHasBlendShapes int `yaml:"legacyComputeAllNormalsFromSmoothingGroupsWhenMeshHasBlendShapes"`
			BlendShapeNormalImportMode                                       int `yaml:"blendShapeNormalImportMode"`
			NormalSmoothingSource                                            int `yaml:"normalSmoothingSource"`
		} `yaml:"tangentSpace"`
		ReferencedClips  []string `yaml:"referencedClips"`
		ImportAnimation  int      `yaml:"importAnimation"`
		HumanDescription struct {
			SerializedVersion int `yaml:"serializedVersion"`
			Human             []struct {
				BoneName  string `yaml:"boneName"`
				HumanName string `yaml:"humanName"`
				Limit     struct {
					Min struct {
						X int `yaml:"x"`
						Y int `yaml:"y"`
						Z int `yaml:"z"`
					} `yaml:"min"`
					Max struct {
						X int `yaml:"x"`
						Y int `yaml:"y"`
						Z int `yaml:"z"`
					} `yaml:"max"`
					Value struct {
						X int `yaml:"x"`
						Y int `yaml:"y"`
						Z int `yaml:"z"`
					} `yaml:"value"`
					Length   int `yaml:"length"`
					Modified int `yaml:"modified"`
				} `yaml:"limit"`
			} `yaml:"human"`
			Skeleton           []Skeleton  `yaml:"skeleton"`
			ArmTwist           float64     `yaml:"armTwist"`
			ForeArmTwist       float64     `yaml:"foreArmTwist"`
			UpperLegTwist      float64     `yaml:"upperLegTwist"`
			LegTwist           float64     `yaml:"legTwist"`
			ArmStretch         float64     `yaml:"armStretch"`
			LegStretch         float64     `yaml:"legStretch"`
			FeetSpacing        int         `yaml:"feetSpacing"`
			GlobalScale        int         `yaml:"globalScale"`
			RootMotionBoneName interface{} `yaml:"rootMotionBoneName"`
			HasTranslationDoF  int         `yaml:"hasTranslationDoF"`
			HasExtraRoot       int         `yaml:"hasExtraRoot"`
			SkeletonHasParents int         `yaml:"skeletonHasParents"`
		} `yaml:"humanDescription"`
		LastHumanDescriptionAvatarSource struct {
			InstanceID int `yaml:"instanceID"`
		} `yaml:"lastHumanDescriptionAvatarSource"`
		AutoGenerateAvatarMappingIfUnspecified  int         `yaml:"autoGenerateAvatarMappingIfUnspecified"`
		AnimationType                           int         `yaml:"animationType"`
		HumanoidOversampling                    int         `yaml:"humanoidOversampling"`
		AvatarSetup                             int         `yaml:"avatarSetup"`
		AddHumanoidExtraRootOnlyWhenUsingAvatar int         `yaml:"addHumanoidExtraRootOnlyWhenUsingAvatar"`
		AdditionalBone                          int         `yaml:"additionalBone"`
		UserData                                interface{} `yaml:"userData"`
		AssetBundleName                         interface{} `yaml:"assetBundleName"`
		AssetBundleVariant                      interface{} `yaml:"assetBundleVariant"`
	} `yaml:"ModelImporter"`
}

type Skeleton struct {
	Name       string      `yaml:"name"`
	ParentName interface{} `yaml:"parentName"`
	Position   struct {
		X float64 `yaml:"x"`
		Y float64 `yaml:"y"`
		Z float64 `yaml:"z"`
	} `yaml:"position"`
	Rotation struct {
		X float64 `yaml:"x"`
		Y float64 `yaml:"y"`
		Z float64 `yaml:"z"`
		W float64 `yaml:"w"`
	} `yaml:"rotation"`
	Scale struct {
		X int `yaml:"x"`
		Y int `yaml:"y"`
		Z int `yaml:"z"`
	} `yaml:"scale"`
}
