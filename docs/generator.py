# This is a really hacky script that scrapes my golang endpoints to generate most of the API documentation. Description and examples need to be added manually.
#
# Please dont use this, its specific to my setup and conventions.

import os
import json

path = "/home/tyler/Development/PasswordManager5/api/endpoints"
apiPrefix = "/api/v1"
dirPathPrefix = "/home/tyler/Development/PasswordManager5/api/endpoints"
outputPath = "/home/tyler/Development/PasswordManager5/docs/api"

errorResponse = {
    "error": {
        "code": "string",
        "message": "string",
    }
}

data = {}

def formatMarkdown(path, entries):
    content = f"""# {path}
If a GET request requires parameters, the parameters can only be passed using query parameters. (GET requests does not support body).

All requests need to have a `Content-type` header that is either `application/json` or `application/xml`, requests wont work otherwise. 

Required/Optional Request/Response parameters are denoted like so:
- `required param` or `<required param>`
- `[optional param]` 

This API is public, however, it is strongly recommended that you use an official client."""
    
    for method, data in entries.items():
        content += f"""

## {method}
### Description
// DESCRIPTION HERE //

### Request Format
```
Headers:
--------
{json.dumps(data["headers"], indent=4) or "None"}


Params:
-------
{json.dumps(data["input"], indent=4) or "None"}
```

### Response Format
```
"""

        for status, responseData in sorted(data["responses"].items(), key=lambda x: int(x[0])):
            underline = ""
            
            if status == 500:
                status = "500 (Server Error)"
            elif status == 400:
                status = "400 (Client Error)"
                
            for i in range(len(f"Status {status}:")):
                underline += "-"
            content += f"""Status {status}:
{underline}
{json.dumps(responseData, indent=4) or "None"}


"""
        content = content[:-3]
        content += f"""
```

### Example
// EXAMPLE HERE //"""

    return content.replace('"..."', "...").replace('"valid authToken"', "valid authToken").replace('"authToken string"', "authToken string").replace('"string"', "string").replace('"url base64 string"', "url base64 string").replace('"base64 string"', "base64 string").replace('"uuid string"', "uuid string").replace('"int"', "int").replace('"bool"', "bool")

        

for (dirPath, dirNames, fileNames) in os.walk(path):
    for fileName in fileNames:
        endpointPath = apiPrefix + (dirPath + "/" + fileName).replace(dirPathPrefix, "")

        splitEndpointPath = endpointPath.split("/")
        folderName = splitEndpointPath[-2:-1][0]
        goFileName = splitEndpointPath[-1:][0].replace(".go", "")

        if folderName == goFileName:
            endpointPath = endpointPath.replace("/" + folderName + ".go", "")
        else:
            endpointPath = endpointPath.replace(".go", "")

        with open(dirPath + "/" + fileName, "r") as fileReader:
            data[endpointPath] = {}
            dataIndex = -1
            dataKey = ""

            finishedStructs = []
            structDataLocation = {}
            structContinue = False
            embededStructDataLocation = {}
            embededStructContinue = False
            embededIsArray = False
            thirdEmbededStructDataLocation = {}
            thirdEmbededStructContinue = False
            thirdEmbededIsArray = False

            for line in fileReader.readlines():
                if line.startswith("func"):
                    dataIndex += 1
                    dataKey = list(data[endpointPath].keys())[dataIndex]
                    
                if line.startswith("type ") and (line.endswith(" struct{}\n") or line.endswith(" struct {\n")):
                    structName = line.replace("type ", "").replace(" struct{}\n", "").replace(" struct {\n", "")

                    method = ""
                    index = 0
                    while structName[index].islower() or index == 0:
                        method += structName[index]
                        index += 1

                    method = method.upper()

                    

                    data[endpointPath][method] = {
                        "input": {},
                        "responses": {},
                        "headers": {},
                    }

                if "Input" in line and line.endswith(" struct {\n"):
                    structContinue = True
                elif "Input" in line and line.endswith(" struct{}\n"):
                    finishedStructs.append({})

                if "c.JSON(400, " in line:
                    data[endpointPath][dataKey]["responses"][400] = errorResponse
                    
                if "c.JSON(500, " in line:
                    data[endpointPath][dataKey]["responses"][500] = errorResponse

                if "authedUser" in line:
                    data[endpointPath][dataKey]["headers"]["Authorization"] = "valid authToken"
                        
                if "c.JSON(200, " in line:
                    response = "{" + line.strip().replace("c.JSON(200, gin.H{", "").replace("})", "") + "}"
                    jsonResponse = response.replace(": ", ': "').replace(", ", '", ').replace("}", '"}').replace('""', '"').replace('" + ', " ").replace('{"}', "{}")
                    responseDict = json.loads(jsonResponse)
                    returnResponseDict = {}

                    for (key, value) in responseDict.items():
                        splitValue = value.split(".")
                        if len(splitValue) == 3 and splitValue[-2:-1][0] == "ID" and splitValue[-1:][0] == "String()":
                            returnResponseDict[key] = "uuid string"
                        elif key == "options":
                            returnResponseDict[key] = {
                                "publicKey": {
                                    "challenge": "url base64 string",
                                    "rp": {
                                        "name": "string",
                                        "[icon]": "string",
                                        "id": "string",
                                    },
                                    "user": {
                                        "name": "string",
                                        "[icon]": "string",
                                        "[displayName]": "string",
                                        "id": "url base64 string",
                                    },
                                    "[pubKeyCredParams]": [
                                        {
                                            "type": "string",
                                            "alg": "int",
                                        },
                                        "...",
                                    ],
                                    "[authenticatorSelection]": {
                                        "[authenticatorAttachment]": "string",
                                        "[requireResidentKey]": "bool",
                                        "[residentKey]": "string",
                                        "[userVerification]": "string",
                                    },
                                    "[timeout]": "int",
                                    "[excludeCredentials]": [
                                        {
                                            "type": "string",
                                            "id": "url base64 string",
                                            "[transports]": ["string", "..."]
                                        },
                                        "...",
                                    ],
                                    "[attestation]": "string",
                                }
                            }
                        elif key == "availableChallenges":
                            returnResponseDict[key] = ["string", "..."]
                        elif key == "webauthnCredentials":
                            returnResponseDict[key] = [
                                {
                                    "id": "uuid string",
                                    "name": "string",
                                    "createdAt": "time",
                                },
                                "...",
                            ]
                        elif key == "webauthnCredential":
                            returnResponseDict[key] = {
                                "id": "uuid string",
                                "name": "string",
                                "createdAt": "time",
                            }
                        elif key == "passwords":
                            returnResponseDict[key] = [
                                {
                                    "id": "uuid string",
                                    "name": "base64 string",
                                    "nameIv": "base64 string",
                                    "username": "base64 string",
                                    "usernameIv": "base64 string",
                                    "password": "base64 string",
                                    "passwordIv": "base64 string",
                                    "colour": "hex colour string",
                                    "additionalFields": [
                                        {
                                            "key": "base64 string",
                                            "keyIv": "base64 string",
                                            "value": "base64 string",
                                            "valueIv": "base64 string",
                                        },
                                    ],
                                    "urls": [
                                        {
                                            "url": "base64 string",
                                            "urlIv": "base64 string",
                                        },
                                    ],
                                },
                                "...",
                            ]
                        elif key == "password":
                            returnResponseDict[key] = {
                                "id": "uuid string",
                                "name": "base64 string",
                                "nameIv": "base64 string",
                                "username": "base64 string",
                                "usernameIv": "base64 string",
                                "password": "base64 string",
                                "passwordIv": "base64 string",
                                "colour": "hex colour string",
                                "additionalFields": [
                                    {
                                        "key": "base64 string",
                                        "keyIv": "base64 string",
                                        "value": "base64 string",
                                        "valueIv": "base64 string",
                                    },
                                ],
                                "urls": [
                                    {
                                        "url": "base64 string",
                                        "urlIv": "base64 string",
                                    },
                                ],
                            }
                        elif key == "authToken":
                            returnResponseDict[key] = "authToken string"
                        else:
                            returnResponseDict[key] = "base64 string"

                    data[endpointPath][dataKey]["responses"][200] = returnResponseDict
                    if len(finishedStructs) > dataIndex:
                        data[endpointPath][dataKey]["input"] = finishedStructs[dataIndex]

                if thirdEmbededStructContinue:
                    stripped = line.strip()
                    if stripped == "}":
                        thirdEmbededStructContinue = False
                    else:
                        parts = line.strip().split(" `")
                        if parts[0] == "}":
                            thirdEmbededStructContinue = False
                            name =  parts[1].replace('form:"', "").replace('" json:"', "//").split("//")[0]
                            if thirdEmbededIsArray:
                                embededStructDataLocation[name] = [thirdEmbededStructDataLocation, "..."]
                            else:
                                embededStructDataLocation[name] = thirdEmbededStructDataLocation
                            thirdEmbededIsArray = False
                            continue
                        elif len(parts) == 2:
                            expectedType = parts[0].replace("  ", " ").replace("  ", " ").replace("  ", " ").replace("  ", " ").replace("  ", " ").replace("  ", " ").replace("  ", " ").replace("  ", " ").split(" ")[1]
                            name = parts[1].replace('form:"', "").replace('" json:"', "//").split("//")[0]
                            thirdEmbededStructDataLocation[name] = expectedType

                if embededStructContinue and not thirdEmbededStructContinue:
                    stripped = line.strip()
                    if stripped == "}":
                        embededStructContinue = False
                    else:
                        parts = line.strip().split(" `")
                        if parts[0] == "}":
                            embededStructContinue = False
                            name =  parts[1].replace('form:"', "").replace('" json:"', "//").split("//")[0]
                            if embededIsArray:
                                structDataLocation[name] = [embededStructDataLocation, "..."]
                            else:
                                structDataLocation[name] = embededStructDataLocation
                            embededIsArray = False
                            continue
                        else:
                            expectedType = parts[0].replace("  ", " ").replace("  ", " ").replace("  ", " ").replace("  ", " ").replace("  ", " ").replace("  ", " ").replace("  ", " ").replace("  ", " ").split(" ")[1]
                            if expectedType == "struct":
                                thirdEmbededStructContinue = True
                                thirdEmbededStructDataLocation = {}
                            elif expectedType == "[]struct":
                                thirdEmbededStructContinue = True
                                thirdEmbededStructDataLocation = {}
                                thirdEmbededIsArray = True
                            elif len(parts) == 2:
                                name = parts[1].replace('form:"', "").replace('" json:"', "//").split("//")[0]
                                embededStructDataLocation[name] = expectedType

                if structContinue and not embededStructContinue and not thirdEmbededStructContinue:
                    stripped = line.strip()
                    if stripped == "}":
                        structContinue = False
                        finishedStructs.append(structDataLocation)
                    else:
                        parts = line.strip().split(" `")
                        if parts[0] == "}":
                            embededStructContinue = False
                        else:
                            expectedType = parts[0].replace("  ", " ").replace("  ", " ").replace("  ", " ").replace("  ", " ").replace("  ", " ").replace("  ", " ").replace("  ", " ").replace("  ", " ").split(" ")[1]
                            if expectedType == "struct":
                                embededStructContinue = True
                                embededStructDataLocation = {}
                            elif expectedType == "[]struct":
                                embededStructContinue = True
                                embededStructDataLocation = {}
                                embededIsArray = True
                            elif len(parts) == 2:
                                name = parts[1].replace('form:"', "").replace('" json:"', "//").split("//")[0]
                                structDataLocation[name] = expectedType

for path, data in data.items():
    markdown = formatMarkdown(path, data)
    fileName = f"{outputPath}{path.replace(apiPrefix, '')}.md"
    os.makedirs(os.path.dirname(fileName), exist_ok=True)
    with open(fileName, "a") as fileWriter:
        fileWriter.write("\n\n" + markdown)

print("Success")
