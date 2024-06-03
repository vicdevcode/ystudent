/*
  Warnings:

  - Added the required column `description` to the `Profile` table without a default value. This is not possible if the table is not empty.
  - Added the required column `fio` to the `Profile` table without a default value. This is not possible if the table is not empty.
  - Added the required column `role` to the `Profile` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "Profile" ADD COLUMN     "description" TEXT NOT NULL,
ADD COLUMN     "fio" TEXT NOT NULL,
ADD COLUMN     "role" TEXT NOT NULL;
