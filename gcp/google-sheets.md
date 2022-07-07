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

