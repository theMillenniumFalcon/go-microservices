import {MigrationInterface, QueryRunner} from "typeorm";

export class Migrations1650629122854 implements MigrationInterface {
    name = 'Migrations1650629122854'

    public async up(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`
            CREATE TABLE "temperature" (
                "id" SERIAL NOT NULL,
                "AmountInches" integer NOT NULL,
                "TempHigh" integer NOT NULL,
                "TempLow" integer NOT NULL,
                "ZipCode" character varying NOT NULL,
                "createdAt" TIMESTAMP NOT NULL DEFAULT now(),
                "updatedAt" TIMESTAMP NOT NULL DEFAULT now(),
                CONSTRAINT "PK_3b69dc45d57daf28f4b930eb4c9" PRIMARY KEY ("id")
            )
        `);
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`
            DROP TABLE "temperature"
        `);
    }

}
