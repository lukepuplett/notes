# Google Sheets API

## Overview

- RESTful API let's you
- Create sheets
- Read and write cell values
- Update formatting
- Manage _Connected Sheets_

### Addressing Examples

- URL scheme is https://docs.google.com/spreadsheets/d/{spreadsheet-id}/edit#gid={sheet-id}

#### A1 notation

- `Sheet1!A1:B2`          = first two cells in top rows
- `Sheet1!A:A`            = cells in first column
- `Sheet1!1:2`            = cells in first two rows
- `Sheet1!A5:A`           = cells in first column from row 5
- `A1:B2`                 = first two cells in first visible sheet
- `Sheet1`                = cells in Sheet1
- `'My Custom Sheet'!A:A` = sheet names with spaces enclosed in `'`
- `'My Custom Sheet'`     = all cells in that sheet

**Note** - Be careful with naming sheets and ranges; `A1` can also mean all cells in a sheet named A1.

#### R1C1 notation

This uses row and columns numbers. Useful when referencing cells relative to a given cell's position.

- `Sheet1!R1C1:R2C2`      = first two cells in rop rows
- `Sheet1!R[3]C[1]`       = three rows below and one right of the current cell

Can also name ranges and protect them from changes.

## Create access credentials

Credentials are needed to get an access token to call **Google Workspace** APIs.

### Credential types

- API key - use this to access publicly-available data anonymously in your app.
- OAuth client ID - use this to authenticate as an end user and access their data.
- Service account - use this to authenticate as a service account, or to access resources on behalf of Google Workspace or Cloud Identity users through domain-wide delegation.

#### API key credentials

This a pseudo-random string used to access documents shared using "Anyone on the Internet with this link". You must create an API key:

- In the GCP console, click **Menu** > **APIs & Services** > **Credentials**
- Choose **Create credentials** > **API key**
- Copy it.

#### OAuth client ID credentials

Your app will need a client ID before you can authenticate as a user. It'll need different client IDs for each platform (web, Android et al).

#### Service account credentials

Once you've created a service account in the GCP console, you'll need to obtain the public/private key pair. You can create these via the console, too, and download them as JSON.

## .NET. quickstart

- Needs a GCP project with Google Sheets API enabled.
- Authorization credentials for a desktop application.
- A Google account.
- `Install-Package Google.Apis.Sheets.v4`
- Unfortunately the sample uses the JSON public/private key.

```
    UserCredential credential;
    
    // Sample uses JSON credential. We need to know how to begin an OAuth flow to
    // get the user's access token.

    // Create Google Sheets API service.
    var service = new SheetsService(new BaseClientService.Initializer
    {
        HttpClientInitializer = credential,
        ApplicationName = ApplicationName
    });

    // Define request parameters.
    String spreadsheetId = "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms";
    String range = "Class Data!A2:E";
    SpreadsheetsResource.ValuesResource.GetRequest request =
        service.Spreadsheets.Values.Get(spreadsheetId, range);

    // Prints the names and majors of students in a sample spreadsheet:
    // https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
    ValueRange response = request.Execute();
    IList<IList<Object>> values = response.Values;
```
##### OAuth 2.0 .NET client library guide

https://developers.google.com/api-client-library/dotnet/guide/aaa_oauth#web-applications-aspnet-mvc  
