create database "InventoryAPI"

create table if not exists "User"(
"id" serial primary key,
"username" varchar unique not null,
"password" varchar not null
);

create table if not exists "Category"(
"id" serial primary key,
"name" varchar not null,
"description" varchar not null
);

create table if not exists "Goods"(
"id" serial primary key,
"category_id" int references "Category"("id") on delete cascade not null,
"name" varchar not null,
"photo_url" varchar not null,
"price" varchar not null,
"purchase_date" date not null,
"total_usage_days" int not null
);

insert into "Category" ("name", "description")
values
  ('Furniture', 'Furniture for office'),
  ('Electronics', 'Office electronics'),
  ('Appliances', 'Appliances for office');
  

insert into "Goods" ("category_id", "name", "photo_url", "price", "purchase_date", "total_usage_days")
values
  (1, 'Office Desk', 'http://example.com/office_desk.jpg', '2500000', '2024-09-25', 45),
  (1, 'Office Chair', 'http://example.com/office_chair.jpg', '1500000', '2024-08-11', 90),
  (2, 'Desktop Computer', 'http://example.com/desktop_computer.jpg', '8000000', '2024-06-11', 120),
  (3, 'Coffee Machine', 'http://example.com/coffee_machine.jpg', '1800000', '2024-10-10', 30),
  (2, 'Laptop', 'http://example.com/laptop_dell_xps_13.jpg', '12000000', '2024-08-26', 75),
  (2, 'Monitor 24 Inch', 'http://example.com/monitor_24inch.jpg', '3500000', '2024-09-25', 45),
  (3, 'Refrigerator', 'http://example.com/office_fridge.jpg', '5000000', '2024-06-30', 100);
  