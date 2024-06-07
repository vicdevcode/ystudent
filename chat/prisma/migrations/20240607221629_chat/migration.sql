/*
  Warnings:

  - A unique constraint covering the columns `[chatId]` on the table `Department` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[chatId]` on the table `Faculty` will be added. If there are existing duplicate values, this will fail.

*/
-- AlterTable
ALTER TABLE "Department" ADD COLUMN     "chatId" TEXT;

-- AlterTable
ALTER TABLE "Faculty" ADD COLUMN     "chatId" TEXT;

-- CreateIndex
CREATE UNIQUE INDEX "Department_chatId_key" ON "Department"("chatId");

-- CreateIndex
CREATE UNIQUE INDEX "Faculty_chatId_key" ON "Faculty"("chatId");

-- AddForeignKey
ALTER TABLE "Faculty" ADD CONSTRAINT "Faculty_chatId_fkey" FOREIGN KEY ("chatId") REFERENCES "Chat"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Department" ADD CONSTRAINT "Department_chatId_fkey" FOREIGN KEY ("chatId") REFERENCES "Chat"("id") ON DELETE SET NULL ON UPDATE CASCADE;
