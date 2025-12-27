# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2024-12-27

### Added
- Initial stable release of the `dg-firebase` plugin.
- **Firebase Auth**: Integration with Firebase Authentication for user management and token verification.
- **FCM (Firebase Cloud Messaging)**: High-performance client for sending push notifications.
- **Firestore**: Integration with Google Cloud Firestore for NoSQL document storage.
- **Observability**: OpenTelemetry instrumentation for FCM and Firestore operations.
- **Container Integration**: Auto-registration with `Injectable` pattern support.

### Features
- Support for multiple Firebase projects/apps.
- Environment-based configuration for credentials.
- Automatic token verification middleware support.
- FCM batch sending and topic management.
- Standardized error handling for Firebase operations.

### Documentation
- Comprehensive README with setup and configuration guides.
- Detailed documentation for FCM, Auth, and Firestore integrations.
- Code examples for common use cases.

---

## Development History

The following milestones represent the journey to v1.0.0:

### 2024-11-24
- Added OpenTelemetry instrumentation for FCM operations.
- Fixed configuration injection bugs in FCM provider.

### 2024-11-23
- Initial implementation of Firestore integration.
- Standardized Firebase App initialization from environment/config.

### 2024-11-22
- Initial implementation of Firebase Auth and basic FCM client.
- Integration with `dg-core` service container.
