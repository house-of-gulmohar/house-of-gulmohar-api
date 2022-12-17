ALTER TABLE product 
ADD COLUMN quantity INT NOT NULL check(quantity >= 0),
ADD COLUMN features BOOLEAN NOT NULL DEFAULT false;