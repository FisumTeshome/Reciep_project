-- Seed data for MinabPro Recipe Application
-- This migration populates the database with initial data

-- Insert default categories
INSERT INTO categories (name, description, color) VALUES
('Breakfast', 'Morning meals to start your day', '#FF6B6B'),
('Lunch', 'Midday meals for energy', '#4ECDC4'),
('Dinner', 'Evening meals and family time', '#45B7D1'),
('Dessert', 'Sweet treats and indulgence', '#96CEB4'),
('Snacks', 'Quick bites and appetizers', '#FFEAA7'),
('Beverages', 'Drinks, smoothies, and cocktails', '#DDA0DD'),
('Appetizers', 'Starters and small plates', '#98D8C8'),
('Soups', 'Warm and cold soups', '#F7DC6F'),
('Salads', 'Fresh and healthy options', '#82E0AA'),
('Bread', 'Fresh baked goods', '#F8C471'),
('Pasta', 'Italian and pasta dishes', '#E74C3C'),
('Seafood', 'Fish and seafood dishes', '#3498DB'),
('Meat', 'Beef, pork, and lamb dishes', '#E67E22'),
('Poultry', 'Chicken and turkey dishes', '#F39C12'),
('Vegetarian', 'Plant-based meals', '#27AE60'),
('Vegan', '100% plant-based meals', '#2ECC71'),
('Gluten-Free', 'Gluten-free options', '#9B59B6'),
('Quick & Easy', '30 minutes or less', '#1ABC9C'),
('Slow Cooker', 'Set it and forget it', '#34495E'),
('One Pot', 'Minimal cleanup meals', '#16A085')
ON CONFLICT (name) DO NOTHING;

-- Insert sample ingredients
INSERT INTO ingredients (name, description, unit) VALUES
('All-purpose flour', 'Basic baking flour', 'cups'),
('Sugar', 'Granulated white sugar', 'cups'),
('Salt', 'Table salt', 'teaspoons'),
('Black pepper', 'Freshly ground black pepper', 'teaspoons'),
('Olive oil', 'Extra virgin olive oil', 'tablespoons'),
('Butter', 'Unsalted butter', 'tablespoons'),
('Eggs', 'Large eggs', 'pieces'),
('Milk', 'Whole milk', 'cups'),
('Onion', 'Yellow onion', 'pieces'),
('Garlic', 'Fresh garlic cloves', 'cloves'),
('Tomatoes', 'Fresh tomatoes', 'pieces'),
('Chicken breast', 'Boneless, skinless chicken breast', 'pounds'),
('Ground beef', 'Lean ground beef', 'pounds'),
('Rice', 'Long grain white rice', 'cups'),
('Pasta', 'Spaghetti or other pasta', 'pounds'),
('Cheese', 'Shredded cheese', 'cups'),
('Lemon', 'Fresh lemon', 'pieces'),
('Lime', 'Fresh lime', 'pieces'),
('Herbs', 'Fresh herbs (basil, parsley, etc.)', 'cups'),
('Spices', 'Various spices', 'teaspoons')
ON CONFLICT (name) DO NOTHING;

-- Create a sample admin user (password: admin123)
-- Note: In production, use proper password hashing
INSERT INTO users (username, email, password, first_name, last_name, is_verified) VALUES
('admin', 'admin@minabpro.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Admin', 'User', true)
ON CONFLICT (username) DO NOTHING; 