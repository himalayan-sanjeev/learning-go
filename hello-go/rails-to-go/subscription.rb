# app/models/subscription.rb
class Subscription < ApplicationRecord
  belongs_to :customer
  has_many :payments, dependent: :restrict_with_error

  enum status: { active: 0, paused: 1, cancelled: 2 }

  validates :customer_id, presence: true
  validates :amount_cents, numericality: { greater_than: 0 }
  validates :currency, presence: true, length: { is: 3 }

  before_validation :normalize_currency
  before_create :set_defaults

  scope :billable, -> { active.where("next_bill_at <= ?", Time.current) }

  def billable?
    active? && next_bill_at.present? && next_bill_at <= Time.current
  end

  def cancel!(reason: nil)
    transaction do
      update!(status: :cancelled, cancelled_at: Time.current, cancel_reason: reason)
      payments.pending.update_all(status: "void")
    end
  end

  private

  def normalize_currency
    self.currency = currency&.upcase
  end

  def set_defaults
    self.status ||= :active
  end
end
