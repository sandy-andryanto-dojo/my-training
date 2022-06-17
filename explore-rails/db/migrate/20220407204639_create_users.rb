class CreateUsers < ActiveRecord::Migration[6.1]
  def change
    create_table :users do |t|
      t.string :username, limit: 191, null: false
      t.string :email, limit: 191, null: false
      t.string :password_digest
      t.boolean :is_active
      t.timestamps
    end
    add_index :users, :email, unique: true
    add_index :users, :username, unique: true
  end
end
