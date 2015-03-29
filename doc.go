/*
   osc provides a package for sending and receiving OpenSoundControl messages.
   The package is implemented in pure Go.

   The implementation is based on the Open Sound Control 1.0 Specification
   (http://opensoundcontrol.org/spec-1_0).

   Open Sound Control (OSC) is an open, transport-independent, message-based
   protocol developed for communication among computers, sound synthesizers,
   and other multimedia devices.

   Features:
   - Supports OSC messages with 'i' (Int32), 'f' (Float32),
    's' (string), 'b' (blob / binary data), 'h' (Int64), 't' (OSC timetag),
     'd' (Double/int64), 'T' (True), 'F' (False), 'N' (Nil) types.
   - OSC bundles, including timetags
   - Support for OSC address pattern including '*', '?', '{,}' and '[]' wildcards
   - TODO: Describe registering methods

   This OSC implementation uses the UDP protocol for sending and receiving
   OSC packets.

   The unit of transmission of OSC is an OSC Packet. Any application that sends
   OSC Packets is an OSC Client; any application that receives OSC Packets is
   an OSC Server.

   An OSC packet consists of its contents, a contiguous block of binary data,
   and its size, the number of 8-bit bytes that comprise the contents. The
   size of an OSC packet is always a multiple of 4.

   OSC packets come in two flavors:

   OSC Messages: An OSC message consists of an OSC address pattern, followed
   by an OSC Type Tag String, and finally by zero or more OSC arguments.

   OSC Bundles: An OSC Bundle consists of the string "#bundle" followed
   by an OSC Time Tag, followed by zero or more OSC bundle elements. Each bundle
   element can be another OSC bundle (note this recursive definition: bundle may
   contain bundles) or OSC message.

   An OSC bundle element consists of its size and its contents. The size is
   an int32 representing the number of 8-bit bytes in the contents, and will
   always be a multiple of 4. The contents are either an OSC Message or an
   OSC Bundle.

   The following argument types are supported: 'i' (Int32), 'f' (Float32),
   's' (string), 'b' (blob / binary data), 'h' (Int64), 't' (OSC timetag),
   'd' (Double/int64), 'T' (True), 'F' (False), 'N' (Nil).

   osc supports the following OSC address patterns:
   - '*', '?', '{,}' and '[]' wildcards.

   Original Author: Sebastian Ruml <sebastian.ruml@gmail.com>
   Created: 2013.08.19

   Cribbed by: Brian Sorahan <bsorahan@gmail.com>
   March 29, 2015
*/
package osc
