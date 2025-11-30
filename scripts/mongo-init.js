// MongoDB initialization script
db = db.getSiblingDB('chat_platform');

// Create collections
db.createCollection('users');
db.createCollection('conversations');
db.createCollection('messages');
db.createCollection('files');

// Create indexes for messages collection
db.messages.createIndex({ "conversation_id": 1, "sequence_number": 1 }, { unique: true });
db.messages.createIndex({ "_id": 1 }, { unique: true });
db.messages.createIndex({ "from": 1 });
db.messages.createIndex({ "to": 1 });
db.messages.createIndex({ "created_at": -1 });

// Create indexes for conversations collection
db.conversations.createIndex({ "_id": 1 }, { unique: true });
db.conversations.createIndex({ "members": 1 });
db.conversations.createIndex({ "last_message_at": -1 });

// Create indexes for files collection
db.files.createIndex({ "_id": 1 }, { unique: true });
db.files.createIndex({ "uploaded_by": 1 });
db.files.createIndex({ "uploaded_at": -1 });

// Create indexes for users collection
db.users.createIndex({ "username": 1 }, { unique: true });

print('MongoDB initialization completed');
