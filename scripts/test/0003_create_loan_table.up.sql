CREATE TABLE IF NOT EXISTS "loan" (
                                      "id" SERIAL PRIMARY KEY,
                                      "user_id" INT NOT NULL,
                                      "book_id" INT NOT NULL,
                                      "loaned_at" TIMESTAMP WITH TIME ZONE NOT NULL,
                                      "due_date" TIMESTAMP WITH TIME ZONE NOT NULL,
                                      "return_date" TIMESTAMP WITH TIME ZONE
);