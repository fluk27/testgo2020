	CREATE TABLE Custromer
	(
	   CustID SERIAL PRIMARY KEY,
	   Custfirstname varchar(255),
	   Custlastname varchar(255),
	   Custusername varchar(255),
	   Custpassword text,
	   statusPDPA  boolean DEFAULT TRUE
	);

	CREATE TABLE Car
	(
	   carID SERIAL PRIMARY KEY,
	   carName varchar(255),
	   carType moo,
	   CustID SERIAL,
	   FOREIGN KEY (CustID) REFERENCES Custromer(CustID)
	);

	CREATE TABLE Expenditure
	(
		expenditureID SERIAL PRIMARY KEY,
		expenditureName VARCHAR(255),
		expenditureType boolean,
		expenditureCount INTEGER,
		expenditurePrice NUMERIC(7,0),
		expenditureDate TIMESTAMP,
		CarID SERIAL,
		FOREIGN KEY (CarID) REFERENCES Car (CarID)

	);
