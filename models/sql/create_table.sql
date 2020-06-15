	CREATE TABLE Custromer
	(
	   CustID SERIAL PRIMARY KEY,
	   Custfirstname varchar(255),
	   Custlastname varchar(255),
	   Custusername varchar(255),
	   Custpassword text,
	   statusPDPA  mood
	);

	CREATE TABLE Car
	(
	   carID SERIAL PRIMARY KEY,
	   carName varchar(255),
	   carType moo,
	   CustID SERIAL,
	   FOREIGN KEY (CustID) REFERENCES Custromer(CustID)
	);