create table UserAccount (
	id SERIAL,
	username VARCHAR(50) not null unique,
	name text not null,
	pin int,
	primary key (id)
);


create table BankAccount (
	id SERIAL,
	name text not null,
	ownerid int,
	primary key (id),
	constraint bankaccountowner
		foreign key (ownerid)
			references UserAccount(id)
);

create table Bucket (
	id SERIAL,
	name text not null,
	ownerid int,
	primary key (id),
	constraint bucketowner
		foreign key (ownerid)
			references UserAccount(id)
);

create table LineItem (
	id SERIAL,
	title text not null,
	description text,
	amount float not null,
	bucket int,
	bank int,
	owner int not null,
	constraint itemowner
		foreign key (owner)
			references UserAccount(id)
);
