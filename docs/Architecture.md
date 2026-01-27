# Architecture

This product use Onion Architecture.

## Project Structure

```
main.go
├─infrastracture
│  ├─external
│  │ ├─notion
│  │ └─qiita
│  └─repository
│    └─filesystem
└─internal
    ├─presentation
    │  ├─cmd
    │  ├─dto
    │  └─args
    ├─domain
    │  ├─entity
    │  ├─valueobject
    │  └─service
    └─application
        ├─repository
        ├─external
        └─usecase
```

## Dependency diagram

```mermaid
graph TD
    subgraph "Domain"
        subgraph "Domain Model"
            Entity[Entity]
            VO[Value Object]
        end
        DS[Domain Service]
    end

    subgraph "Application Usecase"
        UC[Usecase]
        RepoIF[Repository IF]
        ExtIF[ExternalService IF]
    end

    subgraph "Presentation Cobra CLI"
        Cmd[Cobra Command cmd/*]
        Args[Args]
        DTO[DTO ViewModel]
    end

    subgraph "Infrastructure Adapter"
        RepoImpl[RepositoryImpl]
        ExtImpl[ExternalServiceImpl]
        FS[FileSystem]
        NotionAPI[NotionAPI]
        QiitaAPI[QiitaAPI]
    end

    %% dependencies inward
    %% Presentation -> Application
    Cmd --> Args
    Cmd --> UC --> DS --> Entity
    Cmd --> DTO

    %% Application -> Domain
    UC --> DS

    %% Usecase depends on ports
    UC --> RepoIF
    UC --> ExtIF

    %% Infrastructure implements ports
    RepoImpl --> RepoIF --> Entity
    ExtImpl --> ExtIF --> Entity

    %% Infrastructure also depends on Domain types
    RepoImpl --> Entity
    ExtImpl --> Entity

    %% Infrastructure depends on external systems
    RepoImpl --> FS
    ExtImpl --> NotionAPI
    ExtImpl --> QiitaAPI

    %% Domain internal dependencies
    DS --> Entity
    Entity --> VO
```
