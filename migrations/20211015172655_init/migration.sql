-- CreateTable
CREATE TABLE "User" (
    "id" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,
    "name" TEXT NOT NULL,
    "role" INTEGER NOT NULL DEFAULT 0,
    "presence_id" INTEGER NOT NULL,
    "party_id" TEXT NOT NULL,

    CONSTRAINT "User_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Presence" (
    "user_id" TEXT NOT NULL,
    "state" TEXT NOT NULL DEFAULT E'',
    "large_image_key" TEXT NOT NULL DEFAULT E'logo_large',
    "small_image_key" TEXT NOT NULL DEFAULT E'loading',
    "start_time" BIGINT NOT NULL DEFAULT 0,
    "end_time" BIGINT NOT NULL DEFAULT 0,
    "details" TEXT NOT NULL DEFAULT E'',
    "large_image_text" TEXT NOT NULL DEFAULT E'',
    "small_image_text" TEXT NOT NULL DEFAULT E'',
    "party_id" TEXT NOT NULL,

    CONSTRAINT "Presence_pkey" PRIMARY KEY ("user_id")
);

-- CreateTable
CREATE TABLE "Party" (
    "id" TEXT NOT NULL,
    "min" INTEGER NOT NULL DEFAULT 0,
    "max" INTEGER NOT NULL DEFAULT 5,
    "join_secret" TEXT NOT NULL DEFAULT E'',
    "owner_id" TEXT NOT NULL,

    CONSTRAINT "Party_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "Party_owner_id_unique" ON "Party"("owner_id");

-- AddForeignKey
ALTER TABLE "User" ADD CONSTRAINT "User_id_fkey" FOREIGN KEY ("id") REFERENCES "Presence"("user_id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "User" ADD CONSTRAINT "User_party_id_fkey" FOREIGN KEY ("party_id") REFERENCES "Party"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Presence" ADD CONSTRAINT "Presence_party_id_fkey" FOREIGN KEY ("party_id") REFERENCES "Party"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Party" ADD CONSTRAINT "Party_owner_id_fkey" FOREIGN KEY ("owner_id") REFERENCES "User"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
