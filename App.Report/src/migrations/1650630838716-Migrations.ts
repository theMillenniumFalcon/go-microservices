import {MigrationInterface, QueryRunner} from "typeorm";

export class Migrations1650630838716 implements MigrationInterface {
    name = 'Migrations1650630838716'

    public async up(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`
            CREATE TABLE "report" (
                "id" SERIAL NOT NULL,
                "AmountInches" integer NOT NULL,
                "AverageHigh" integer NOT NULL,
                "AverageLow" integer NOT NULL,
                "RainfallTotalInches" integer NOT NULL,
                "SnowfallTotalInches" integer NOT NULL,
                "ZipCode" character varying NOT NULL,
                "createdAt" TIMESTAMP NOT NULL DEFAULT now(),
                "updatedAt" TIMESTAMP NOT NULL DEFAULT now(),
                CONSTRAINT "PK_99e4d0bea58cba73c57f935a546" PRIMARY KEY ("id")
            )
        `);
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`
            DROP TABLE "report"
        `);
    }

}
