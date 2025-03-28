

# **CS2 Items API DB Populator**  

## **Setup**  
### **Environment Variables**  
Create a `.env` file with the following format:  

```ini
DATABASE_URL=your_database_url_here
```

## **Usage**  
Run the following command to create and populate the CS2 items database:  

```sh
go run main.go
```

## **Description**  
This script populates a database with all **CS2** items, including:  
- **Skins**: Skin templates for weapons.  
- **Skin Items**:  Representations of how a skin is applied to a weapon.  
- **Keychains**:   Representation of Keychains.
- **Stickers**:    Representation of Stickers.
- **Cases**:       Representation of Cases.
- **Collections**: All CS2 Collections.

## **Notes**  
- **Skins**: A **Skin** is a template applicable to multiple items (ex. Field Tested, Factory New version, etc. of a given skin).  
- **Skin Items**: A **Skin Item** is a specific variation of a **Skin**.  
- **Items**: An **Item** represents an actual entity in the game economy.
