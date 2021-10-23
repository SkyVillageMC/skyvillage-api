/*
  Warnings:

  - You are about to drop the column `owner_id` on the `Party` table. All the data in the column will be lost.

*/
-- DropForeignKey
ALTER TABLE "Party" DROP CONSTRAINT "Party_owner_id_fkey";

-- DropIndex
DROP INDEX "Party_owner_id_unique";

-- AlterTable
ALTER TABLE "Party" DROP COLUMN "owner_id";
