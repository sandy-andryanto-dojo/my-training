class CreateContacts < ActiveRecord::Migration[6.1]
  def change
    create_table :contacts do |t|
      t.string :name, limit: 191, null: false
      t.string :email, limit: 191, null: true
      t.string :phone, limit: 191, null: true
      t.string :website, limit: 191, null: true
      t.text :address, limit: 191, null: true
      t.timestamps
    end
  end
end
