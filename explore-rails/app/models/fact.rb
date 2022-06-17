class Fact < ApplicationRecord
  audited
  belongs_to :user
end
