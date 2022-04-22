import { Entity, PrimaryGeneratedColumn, Column, CreateDateColumn, UpdateDateColumn, BaseEntity } from "typeorm";

@Entity()
export class Temperature extends BaseEntity {
  @PrimaryGeneratedColumn()
  id!: number;

  @Column()
  AmountInches!: number;

  @Column()
  TempHigh!: number;

  @Column()
  TempLow!: number;

  @Column()
  ZipCode!: string;

  @CreateDateColumn()
  createdAt: Date;

  @UpdateDateColumn()
  updatedAt: Date;
}