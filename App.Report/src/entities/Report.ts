import { Entity, PrimaryGeneratedColumn, Column, CreateDateColumn, UpdateDateColumn, BaseEntity } from "typeorm";

@Entity()
export class Report extends BaseEntity {
  @PrimaryGeneratedColumn()
  id!: number;

  @Column()
  AmountInches!: number;

  @Column()
  AverageHigh!: number;

  @Column()
  AverageLow!: number;

  @Column()
  RainfallTotalInches!: number;

  @Column()
  SnowfallTotalInches!: number;

  @Column()
  ZipCode!: string;

  @CreateDateColumn()
  createdAt: Date;

  @UpdateDateColumn()
  updatedAt: Date;
}