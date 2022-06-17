class CreateFacts < ActiveRecord::Migration[6.1]
  def change
    create_table :facts do |t|
      t.references :user, null: false, foreign_key: true
      t.string :fact, limit: 191, null: false
      t.integer :likes, null: true

      t.timestamps
    end
  end
end
