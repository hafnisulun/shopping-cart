INSERT INTO "products" ("uuid", "created_at", "updated_at", "sku", "name", "price", "qty") VALUES
('c1b6214c-1c6c-489e-b763-d195903f5bb5', NOW(), NOW(), '120P90', 'Google Home', 49.99, 10),
('f126e51a-d9f7-4538-b670-c0132eb96b01', NOW(), NOW(), '43N23P', 'MacBook Pro', 5399.99, 5),
('6bdd17e6-3ed9-4ccd-a69a-1afe1883beca', NOW(), NOW(), 'A304SD', 'Alexa Speaker', 109.50, 10),
('ff11dc2a-c15c-4720-8cb8-abf74f38ccf8', NOW(), NOW(), '234234', 'Raspberry Pi B', 30.00, 2);

INSERT INTO "promos" ("uuid", "created_at", "updated_at", "type") VALUES
('d600dbc7-8088-4078-b833-8085303ffbc3', NOW(), NOW(), 'buy_a_get_b'),
('41d1ea20-7c41-41d5-818a-86dbf8cbd8d2', NOW(), NOW(), 'buy_a_get_b'),
('a2a0f9ad-14e3-4073-8e58-6949ae5a8408', NOW(), NOW(), 'min_qty');

INSERT INTO "buy_a_get_b_promos" ("uuid", "created_at", "updated_at", "promo_id", "buy_product_id", "buy_qty", "get_product_id", "get_qty") VALUES
('80c0ce6a-5813-4070-a91b-6c3c78303dff', NOW(), NOW(), 1, 2, 1, 4, 1),
('29216d5e-4431-49e1-a3e3-b32a044a9ed8', NOW(), NOW(), 2, 1, 2, 1, 1);

INSERT INTO "min_qty_promos" ("uuid", "created_at", "updated_at", "promo_id", "product_id", "qty", "discount") VALUES
('0da03145-6fbe-4ee0-b3d6-1d9d3eed2e17', NOW(), NOW(), 3, 3, 3, 10);

INSERT INTO "carts" ("uuid", "created_at", "updated_at") VALUES
('2c699ef7-35cb-4b5a-b769-3224fd69402b', NOW(), NOW()),
('7907fac3-8847-4744-bf83-fb1fe97f6a29', NOW(), NOW()),
('3459341d-a15e-4266-a21d-dec0a337468d', NOW(), NOW());

INSERT INTO "cart_items" ("uuid", "created_at", "updated_at", "cart_id", "product_id", "qty") VALUES
('846cfa87-38d1-487e-9f8f-2edbfda346ad', NOW(), NOW(), 1, 2, 1),
('28031212-7f71-4669-ad9e-d75e598aea47', NOW(), NOW(), 1, 4, 1),
('fef30c0d-1230-4765-a873-681ed1100993', NOW(), NOW(), 2, 1, 3),
('ce6d44fd-5a0b-46a9-8a7d-d985c8870ce2', NOW(), NOW(), 3, 3, 3);
