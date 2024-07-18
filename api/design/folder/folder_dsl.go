package folder

import (
	. "goa.design/goa/v3/dsl"
	. "mai.today/api/design/dsl"
)

var ServiceName = Service("folder", func() {
	Description("The folder service provides operations for managing folders.")
	ErrorInvalidToken()
	createFolder("CreateFolder")
	deleteFolder("DeleteFolder")
	updateFolder("UpdateFolder")
	listFolders("ListFolders")
}).Name

var createFolderResult = CommandResult("createFolder", CreateFolderResultData)

var CreateFolderResultData = ResultType("application/vnd.createfolderresultdata+json", func() {
	Description("Create Folder Result Data")
	Attributes(func() {
		Attribute("folderId", String, "Folder ID", func() {
			Description("Folder identifier")
			Example("123456")
		})
		Attribute("baseId", String, "Base ID", func() {
			Description("Base identifier")
			Example("base123")
		})
		Attribute("parentId", String, "Parent ID", func() {
			Description("Parent identifier")
			Example("parent123")
		})
		Attribute("position", Float64, "Folder position", func() {
			Description("Position of the folder")
			Example(1.0)
		})
		Attribute("createdAt", String, "Creation time", func() {
			Description("Time of creation")
			Example("2024-01-01T00:00:00Z")
		})
		Attribute("updatedAt", String, "Update time", func() {
			Description("Time of update")
			Example("2024-01-02T00:00:00Z")
		})
		Attribute("data", FolderData, "Custom folder data", func() {
			Description("Data associated with the folder")
		})
		Attribute("type", String, "Folder type", func() {
			Description("Type of the folder")
			Example("type123")
		})
		Required("folderId", "baseId", "parentId", "createdAt", "updatedAt", "data", "type")
	})
	View("default", func() {
		Attribute("folderId")
		Attribute("baseId")
		Attribute("parentId")
		Attribute("position")
		Attribute("createdAt")
		Attribute("updatedAt")
		Attribute("data")
		Attribute("type")
	})
})

var deleteFolderResult = CommandResult("deleteFolder", CreateFolderResultData)

var updateFolderResult = CommandResult("updateFolder", CreateFolderResultData)

var listFoldersResult = CommandResult("listFolders", ArrayOf(CreateFolderResultData))

func createFolder(methodName string) {
	const sum = "Create a new folder"
	var resType = createFolderResult
	AsyncMethod(methodName, sum, resType)

	Method(methodName, func() {
		MetaSummary(sum)
		Payload(func() {
			AttributeJWT()
			Attribute("folder", Folder)
			Required("folder")
		})
		Result(resType)
		HTTP(func() {
			POST("/folders")
			ResponseDefaults()
			ResponseBadRequest()
		})
	})
}

func deleteFolder(methodName string) {
	const sum = "Delete a folder by ID"
	var resType = deleteFolderResult
	AsyncMethod(methodName, sum, resType)

	Method(methodName, func() {
		MetaSummary(sum)
		Payload(func() {
			AttributeJWT()
			Attribute("folderId", String, "Folder ID", func() {
				Description("Folder identifier")
				Example("123456")
			})
			Required("folderId")
		})
		Result(resType)
		HTTP(func() {
			DELETE("/folders/{folderId}")
			ResponseDefaults()
		})
	})
}

func updateFolder(methodName string) {
	const sum = "Update folder information by ID"
	var resType = updateFolderResult
	AsyncMethod(methodName, sum, resType)

	Method(methodName, func() {
		MetaSummary(sum)
		Payload(func() {
			AttributeJWT()
			Attribute("folderId", String, "Folder ID", func() {
				Description("Folder identifier")
				Example("123456")
			})
			Attribute("folder", Folder, "Folder data to update")
			Required("folderId", "folder")
		})
		Result(resType)
		HTTP(func() {
			PUT("/folders/{folderId}")
			ResponseDefaults()
			ResponseBadRequest()
		})
	})
}

func listFolders(methodName string) {
	const sum = "List all folders"
	var resType = listFoldersResult
	AsyncMethod(methodName, sum, resType)

	Method(methodName, func() {
		MetaSummary(sum)
		Payload(func() {
			AttributeJWT()
		})
		Result(resType)
		HTTP(func() {
			GET("/folders")
			ResponseDefaults()
			ResponseBadRequest()
		})
	})
}

var Folder = Type("Folder", func() {
	Description("Folder information.")
	Attribute("folderId", String, "Folder ID", func() {
		Description("Folder identifier")
		Example("123456")
	})
	Attribute("baseId", String, "Base ID", func() {
		Description("Base identifier")
		Example("base123")
	})
	Attribute("parentId", String, "Parent ID", func() {
		Description("Parent identifier")
		Example("parent123")
	})
	Attribute("position", Float64, "Folder position", func() {
		Description("Position of the folder")
		Example(1.0)
	})
	Attribute("createdAt", String, "Creation time", func() {
		Description("Time of creation")
		Example("2024-01-01T00:00:00Z")
	})
	Attribute("updatedAt", String, "Update time", func() {
		Description("Time of update")
		Example("2024-01-02T00:00:00Z")
	})
	Attribute("data", FolderData, "Custom folder data", func() {
		Description("Data associated with the folder")
	})
	Attribute("type", String, "Folder type", func() {
		Description("Type of the folder")
		Example("type123")
	})
	Required("folderId", "baseId", "parentId", "createdAt", "updatedAt", "data", "type")
})

var FolderResult = ResultType("application/vnd.folder+json", func() {
	Description("Folder result type")
	Attributes(func() {
		Attribute("folderId", String, "Folder ID", func() {
			Description("Folder identifier")
			Example("123456")
		})
		Attribute("baseId", String, "Base ID", func() {
			Description("Base identifier")
			Example("base123")
		})
		Attribute("parentId", String, "Parent ID", func() {
			Description("Parent identifier")
			Example("parent123")
		})
		Attribute("position", Float64, "Folder position", func() {
			Description("Position of the folder")
			Example(1.0)
		})
		Attribute("createdAt", String, "Creation time", func() {
			Description("Time of creation")
			Example("2024-01-01T00:00:00Z")
		})
		Attribute("updatedAt", String, "Update time", func() {
			Description("Time of update")
			Example("2024-01-02T00:00:00Z")
		})
		Attribute("data", FolderData, "Custom folder data", func() {
			Description("Data associated with the folder")
		})
		Attribute("type", String, "Folder type", func() {
			Description("Type of the folder")
			Example("type123")
		})
		Required("folderId", "baseId", "parentId", "createdAt", "updatedAt", "data", "type")
	})
	View("default", func() {
		Attribute("folderId")
		Attribute("baseId")
		Attribute("parentId")
		Attribute("position")
		Attribute("createdAt")
		Attribute("updatedAt")
		Attribute("data")
		Attribute("type")
	})
})

var FolderData = Type("FolderData", func() {
	Description("Custom folder data.")
	Attribute("color", String, "Folder color", func() {
		Example("red")
	})
	Attribute("name", String, "Folder name", func() {
		Example("My Folder")
	})
	Required("color", "name")
})
