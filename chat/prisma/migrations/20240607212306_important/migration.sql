-- AlterTable
ALTER TABLE "Message" ADD COLUMN     "is_news" BOOLEAN NOT NULL DEFAULT false,
ADD COLUMN     "is_task" BOOLEAN NOT NULL DEFAULT false,
ALTER COLUMN "important" SET DEFAULT false;
