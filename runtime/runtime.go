package runtime

type OwningType struct
{
  Owner bool
}

type UniqueLifecycle interface
{
  IsOwner() bool
  Destructor()
}

