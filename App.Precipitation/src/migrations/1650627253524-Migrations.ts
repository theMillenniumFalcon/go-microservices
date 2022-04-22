import {MigrationInterface, QueryRunner} from "typeorm";

export class Migrations1650627253524 implements MigrationInterface {
    name = 'Migrations1650627253524'

    public async up(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`
            CREATE TABLE "precipitation" (
                "id" SERIAL NOT NULL,
                "AmountInches" integer NOT NULL,
                "WeatherType" character varying NOT NULL,
                "ZipCode" character varying NOT NULL,
                "createdAt" TIMESTAMP NOT NULL DEFAULT now(),
                "updatedAt" TIMESTAMP NOT NULL DEFAULT now(),
                CONSTRAINT "PK_4b2b18da48e8d835622addd241a" PRIMARY KEY ("id")
            )
        `);
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`
            DROP TABLE "precipitation"
        `);
    }

}
