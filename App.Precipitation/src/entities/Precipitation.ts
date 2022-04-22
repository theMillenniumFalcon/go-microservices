import { Entity, PrimaryGeneratedColumn, Column, CreateDateColumn, UpdateDateColumn, BaseEntity } from "typeorm";

@Entity()
export class Precipitation extends BaseEntity {
  @PrimaryGeneratedColumn()
  id!: number;

  @Column()
  AmountInches!: number;

  @Column()
  WeatherType!: string;

  @Column()
  ZipCode!: string;

  @CreateDateColumn()
  createdAt: Date;

  @UpdateDateColumn()
  updatedAt: Date;
}